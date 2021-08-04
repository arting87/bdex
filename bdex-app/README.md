# exchange-view
> exchange front 

## Table of contents

[TOC]

## Requirements

> ###### Installation

- Node.js ：[Downloads](https://nodejs.org/en/download/)
- Yarn@1.15.2：[Downloads](https://yarnpkg.com/zh-Hans/docs/install)
- Webpack

## Dependencies

- Vue + Vue-router + Vant
- axios

## Build Setup

```bash
# install dependencies
yarn install

# serve with hot reload at localhost:8081
yarn run dev

# build for production with minification
yarn run build

# build for production and view the bundle analyzer report
yarn run build --report

# run unit tests
yarn run unit

# run e2e tests
yarn run e2e

# run all tests
yarn test
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).


```bash
element-ui:            
            yarn add element-ui -S
            yarn remove
完整引入：
import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';

Vue.use(ElementUI);

new Vue({
  el: '#app',
  render: h => h(App)
});
按需引入：
```

1. 在lib/eos/eosTrade.js里替换部署eos合约账户名
2. 在lib/eth/exchange.js里替换部署eth合约账户名
