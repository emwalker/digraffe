import 'core-js'
import 'regenerator-runtime/runtime'
import React from 'react'
import BrowserProtocol from 'farce/lib/BrowserProtocol'
import createInitialFarceRouter from 'found/lib/createInitialFarceRouter'
import { hydrate } from 'react-dom'
import { createStore } from 'redux'
import { Provider } from 'react-redux'
import '@primer/css/index.scss?global'

import ClientFetcher from './ClientFetcher'
import {
  createResolver,
  createRouteConfig,
  historyMiddlewares,
  render,
} from '../router'
import './global.scss'

/* eslint no-underscore-dangle: 0 */

const reducer = store => store

const init = async () => {
  const fetcher = new ClientFetcher(window.__RELAY_PAYLOADS__)
  delete window.__RELAY_PAYLOADS__
  const resolver = createResolver(fetcher)

  const preloadedState = window.__PRELOADED_STATE__
  delete window.__PRELOADED_STATE__
  const store = createStore(reducer, preloadedState)
  const routeConfig = createRouteConfig(store)

  const Router = await createInitialFarceRouter({
    historyProtocol: new BrowserProtocol(),
    historyMiddlewares,
    resolver,
    render,
    routeConfig,
  })

  hydrate(
    <Provider store={store}>
      <Router resolver={resolver} />
    </Provider>,
    document.getElementById('root'),
  )

  if (module.hot) module.hot.accept()
}

init()