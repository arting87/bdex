var path = require('path')
var webpack = require('webpack')
var config = require('../config')
var utils = require('./utils')
var projectRoot = path.resolve(__dirname, '../')

module.exports = {
  entry: {
    app: './src/main.js'
  },
  output: {
    path: config.build.assetsRoot,
    publicPath: process.env.NODE_ENV === 'production' ? config.build.assetsPublicPath : config.dev.assetsPublicPath,
    filename: '[name].js'
  },
  resolve: {
    extensions: ['', '.js', '.vue'],
    fallback: [path.join(__dirname, '../node_modules')],
    alias: {
      'vue$': 'vue/dist/vue.js',
      'src': path.resolve(__dirname, '../src'),
      'assets': path.resolve(__dirname, '../src/assets'),
      'components': path.resolve(__dirname, '../src/components'),
      'jquery': path.resolve(__dirname, '../src/js/jquery.js'),
      'moment':path.resolve(__dirname, '../src/plugins/daterangepicker/moment.js'),
      'iCheck':path.resolve(__dirname, '../src/plugins/iCheck/icheck.min.js')
    }
  },
  plugins:[
    new webpack.ProvidePlugin({
      'moment':'moment',
      $:"jquery",
      jQuery:"jquery",
      "window.jQuery":"jquery",
      
      'iCheck':'iCheck'
    })
  ],
  resolveLoader: {
    fallback: [path.join(__dirname, '../node_modules')]
  },
  module: {
    loaders: [
      {
        test: /\.vue$/,
        loader: 'vue'
      },
      {
        test: /\.js$/,
        loader: 'babel',
        include: projectRoot,
        exclude: /(node_modules|plugins)/
      },
      // {  
      //   test: /\.css$/,  
      //   include: [  
      //    /src/,//表示在src目录下的css需要编译  
      //    '/node_modules/element-ui/lib/'   //增加此项  
      //     ],  
      //     loader:'style-loader!css-loader'  
      // }, 
      {
        test: /\.json$/,
        loader: 'json'
      },
      {
        test: /\.css$/,
        include: [  
          /src/,//表示在src目录下的css需要编译  
          '/node_modules/element-ui/lib/'   //增加此项  
        ],  
        loader:'style-loader!css-loader'
      },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        loader: 'url',
        query: {
          limit: 10000,
          name: utils.assetsPath('img/[name].[hash:7].[ext]')
        }
      },
      {
        test: /\.less$/, 
        loader: 'style!css!less'
      },
      {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
        loader: 'url',
        query: {
          limit: 10000,
          name: utils.assetsPath('fonts/[name].[hash:7].[ext]')
        }
      }
    ]
  },
  vue: {
    loaders: utils.cssLoaders(),
    postcss: [
      require('autoprefixer')({
        browsers: ['last 2 versions']
      })
    ]
  }
}
