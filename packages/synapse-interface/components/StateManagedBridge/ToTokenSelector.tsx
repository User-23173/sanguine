import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeShowBorder, useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()
  const { fromToken, toToken } = useBridgeState()
  const { showToTokenBorder } = useBridgeShowBorder()
  const BASE_BUTTON_PROPERTIES =
    'p-md rounded-sm min-w-[80px] bg-[#565058] border'

  let buttonContent
  let buttonClassName

  if (toToken) {
    const src = toToken?.icon?.src
    const symbol = toToken?.symbol

    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      border-transparent
      ${getMenuItemHoverBgForCoin(toToken?.color)}
      ${getBorderStyleForCoinHover(toToken?.color)}
    `

    buttonContent = (
      <div className="flex items-center space-x-2">
        <div className="flex-none hidden md:inline-block">
          <img src={src} alt={symbol} className="w-6 h-6" />
        </div>
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{symbol}</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  } else {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      ${
        showToTokenBorder
          ? 'border-synapsePurple hover:border-secondary'
          : 'border-transparent hover:border-secondary'
      }
    `
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">Out</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-destination-token"
      className={buttonClassName}
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
