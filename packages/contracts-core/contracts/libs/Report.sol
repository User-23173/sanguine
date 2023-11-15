// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";
import { Attestation } from "./Attestation.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @notice Library for formatting the Guard Reports.
 * Reports are submitted to Home contracts in order to slash a fraudulent Notary.
 * Reports are submitted to ReplicaManager contracts in order to blacklist
 * an allegedly fraudulent Notary.
 * Just like an Attestation, a Report could be checked on Home contract
 * on the chain the reported Notary is attesting.
 * Report includes:
 * - Flag, indicating whether the reported attestation is fraudulent.
 * - Reported Attestation (Attestation data and Notary signature on that data).
 * - Guard signature on Report data.
 */
library Report {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev More flag values could be added in the future,
     *      e.g. flag indicating "type" of fraud.
     *      Going forward, Flag.Valid is guaranteed to be
     *      the only Flag specifying a valid attestation.
     *
     *      Flag.Valid indicates a reported valid Attestation.
     *      Flag.Fraud indicates a reported fraud Attestation.
     */
    enum Flag {
        Valid,
        Fraud
    }

    /**
     * @dev ReportData memory layout
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 041): attData        bytes   40 bytes
     *
     * guardSig is Guard's signature on ReportData
     *
     *      Report memory layout
     * [000 .. 002): attLength      uint16   2 bytes (length == AAA - 3)
     * [002 .. 003): flag           uint8    1 bytes
     * [003 .. AAA]: attestation    bytes   ?? bytes (40 + 64/65 bytes)
     * [AAA .. END): guardSig       bytes   ?? bytes (64/65 bytes)
     *
     *      Unpack attestation field (see Attestation.sol)
     * [000 .. 002): attLength      uint16   2 bytes (length == AAA - 3)
     * [002 .. 003): flag           uint8    1 bytes
     * [003 .. 043]: attData        bytes   40 bytes
     * [043 .. AAA): notarySig      bytes   ?? bytes (64/65 bytes)
     * [AAA .. END): guardSig       bytes   ?? bytes (64/65 bytes)
     *
     * notarySig is Notary's signature on AttestationData
     *
     * flag + attData = reportData (see above), so
     *
     *      Report memory layout (sliced alternatively)
     * [000 .. 002): attLength      uint16   2 bytes (length == AAA - 3)
     * [002 .. 043): reportData     bytes   41 bytes
     * [043 .. AAA): notarySig      bytes   ?? bytes (64/65 bytes)
     * [AAA .. END): guardSig       bytes   ?? bytes (64/65 bytes)
     */

    uint256 internal constant OFFSET_ATTESTATION_LENGTH = 0;
    uint256 internal constant OFFSET_FLAG = 2;
    uint256 internal constant OFFSET_ATTESTATION = 3;

    uint256 internal constant ATTESTATION_DATA_LENGTH = 40;
    uint256 internal constant REPORT_DATA_LENGTH = 1 + ATTESTATION_DATA_LENGTH;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    modifier onlyReport(bytes29 _view) {
        _view.assertType(SynapseTypes.REPORT);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       FORMATTERS: REPORT DATA                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns formatted report data with provided fields
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatReportData(Flag _flag, bytes memory _attestation)
        internal
        view
        returns (bytes memory)
    {
        // Extract attestation data from payload
        bytes memory attestationData = _attestation.castToAttestation().attestationData().clone();
        // Construct report data
        return abi.encodePacked(uint8(_flag), attestationData);
    }

    /**
     * @notice Returns formatted report data on valid attestation with provided fields
     * @param _validAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatValidReportData(bytes memory _validAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Valid, _validAttestation);
    }

    /**
     * @notice Returns formatted report data on fraud attestation with provided fields
     * @param _fraudAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatFraudReportData(bytes memory _fraudAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Fraud, _fraudAttestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          FORMATTERS: REPORT                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a properly typed bytes29 pointer for a report payload.
     */
    function castToReport(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.REPORT);
    }

    /**
     * @notice Returns formatted report payload with provided fields.
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @param _guardSig     Guard signature on reportData (see formatReportData below)
     * @return Formatted report
     **/
    function formatReport(
        Flag _flag,
        bytes memory _attestation,
        bytes memory _guardSig
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(uint16(_attestation.length), uint8(_flag), _attestation, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a valid attestation with provided fields.
     * @param _validAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatValidReport(bytes memory _validAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Valid, _validAttestation, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a fraud attestation with provided fields.
     * @param _fraudAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatFraudReport(bytes memory _fraudAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Fraud, _fraudAttestation, _guardSig);
    }

    /**
     * @notice Checks that a payload is a formatted Report payload.
     */
    function isReport(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // Attestation length & flag should exist
        if (length < OFFSET_ATTESTATION) return false;
        uint256 attestationLength = _attestationLength(_view);
        // Guard signature needs to exist
        if (length <= OFFSET_ATTESTATION + attestationLength) return false;
        // Attestation needs to be formatted as well
        return reportedAttestation(_view).isAttestation();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            REPORT SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns whether Report's Flag is Fraud (indicating fraudulent attestation).
     */
    function reportedFraud(bytes29 _view) internal pure onlyReport(_view) returns (bool) {
        return _view.indexUint(OFFSET_FLAG, 1) != uint8(Flag.Valid);
    }

    /**
     * @notice Returns Report's Attestation (which is supposed to be signed by the Notary already).
     */
    function reportedAttestation(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        return _view.slice(OFFSET_ATTESTATION, _attestationLength(_view), SynapseTypes.ATTESTATION);
    }

    /**
     * @notice Returns Report's Data, that is going to be signed by the Guard.
     */
    function reportData(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        // reportData starts from Flag
        return _view.slice(OFFSET_FLAG, REPORT_DATA_LENGTH, SynapseTypes.REPORT_DATA);
    }

    /**
     * @notice Returns Guard's signature on ReportData.
     */
    function guardSignature(bytes29 _view) internal pure onlyReport(_view) returns (bytes29) {
        uint256 offsetSignature = OFFSET_ATTESTATION + _attestationLength(_view);
        return _view.slice(offsetSignature, _view.len() - offsetSignature, SynapseTypes.SIGNATURE);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev No type checks in private functions,
     *      as the type is checked in the function that called this one.
     */
    function _attestationLength(bytes29 _view) private pure returns (uint256) {
        return _view.indexUint(OFFSET_ATTESTATION_LENGTH, 2);
    }
}