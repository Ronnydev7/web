import type {Configuration} from 'webpack';

import path from 'path';

import HtmlWebpackPlugin from 'html-webpack-plugin';

import 'webpack-dev-server';

function localConfig(): Configuration {
  return {
    mode: 'development',
    output: {
      publicPath: '/',
    },
    entry: path.join(__dirname, 'src', 'rootEntry.tsx'),
    module: {
      rules: [
        {
          test: /\.tsx?$/,
          exclude: /^(node_modules|server|__tests__|__test_utils__)$/,
          use: {
            loader: 'babel-loader',
            options: {
              presets: [
                '@babel/preset-env',
                '@babel/preset-react',
                '@babel/preset-typescript',
              ],
            },
          },
        },
      ],
    },
    resolve: {
      extensions: ['.tsx', '.ts', '.js'],
    },
    devServer: {
      static: {
        directory: path.join(__dirname, 'dist'),
        watch: {
          ignored: [
            /\/__tests__\//,
          ],
        },
      },
      compress: false,
      historyApiFallback: true,
      hot: true,
      host: '0.0.0.0',
      port: 4000,
      allowedHosts: [],
    },
    devtool: 'inline-source-map',
    plugins: [
      new HtmlWebpackPlugin({
        template: 'src/index.html',
      }),
    ],
  };
}

export default localConfig;