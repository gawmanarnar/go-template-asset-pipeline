const webpack = require('webpack');
const glob = require('glob');
const path = require('path');

const CopyWebpackPlugin = require('copy-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const ManifestPlugin = require('webpack-manifest-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');

const PROD = process.env.NODE_ENV || 'development';

var entries = {
  styles: ['./assets/css/styles.scss']
};

glob.sync('./assets/*/*.*').reduce((_, entry) => {
  let key = entry.replace(/(\.\/assets\/(js|css|go)\/)|\.(js|s[ac]ss|go)/g, '');
  if (key.startsWith('_') || /(js|s[ac]ss|go)$/i.test(entry) == false) {
    return;
  }

  if (entries[key] == null) {
    entries[key] = [entry];
    return;
  }

  entries[key].push(entry);
});

module.exports = {
  entry: entries,
  output: {
    filename: '[name].[chunkhash].js',
    path: path.resolve(__dirname, 'dist')
  },
  plugins: [
    new CleanWebpackPlugin(['dist'], { verbose: false }),
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery'
    }),
    new MiniCssExtractPlugin({ filename: '[name].[chunkhash].css' }),
    new CopyWebpackPlugin(
      [
        {
          from: './assets',
          to: ''
        }
      ],
      {
        copyUnmodified: true,
        ignore: ['css/**', 'js/**']
      }
    ),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false
    }),
    new ManifestPlugin({
      fileName: 'manifest.json'
    })
  ],
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/
      },
      {
        test: /\.s[ac]ss$/,
        use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader']
      },
      {
        test: /\.(woff|woff2|ttf|svg)(\?v=\d+\.\d+\.\d+)?$/,
        use: 'url-loader'
      },
      { test: /\.eot(\?v=\d+\.\d+\.\d+)?$/, use: 'file-loader' },
      {
        test: require.resolve('jquery'),
        use: 'expose-loader?jQuery!expose-loader?$'
      }
    ]
  }
};

if (PROD != 'development') {
  module.exports.plugins.push(
    new webpack.optimize.UglifyJsPlugin({
      beautify: false,
      mangle: {
        screw_ie8: true,
        keep_fnames: true
      },
      compress: {
        screw_ie8: true
      },
      comments: false
    })
  );
}
