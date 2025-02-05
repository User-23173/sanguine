const path = require('path')

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: false,
  // open source project, will not affect performance (one extra comment ot parse)
  productionBrowserSourceMaps: true,
  webpack: (config, { isServer }) => {
    config.module.rules.push({
      test: /\.nvmrc$/,
      use: [
        {
          loader: path.resolve(__dirname, './nvmrc-loader.js'),
        },
      ],
    })
    config.resolve.extensions = [
      '.js',
      '.jsx',
      '.ts',
      '.tsx',
      '.json',
      '.nvmrc',
    ]
    return config
  },
  eslint: {
    // TODO: enable
    ignoreDuringBuilds: true,
  },
  typescript: {
    tsconfigPath: './tsconfig.json',
  },
}

module.exports = {
  ...nextConfig,
}
