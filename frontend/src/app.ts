import m from 'mithril'
import './i18n'
import RequireAuth from './auth/RequireAuth'

// views
import signIn from './pages/signIn'
import Dashboard from './pages/dashboard'
import Products from './pages/products'
import Product from './pages/product'
import Coupons from './pages/coupons'
import CreateProduct from './pages/createProduct'
import Orders from './pages/orders'
import Order from './pages/order'

// layout
import Layout from './components/layout'

m.route(document.body, '/signin', {
  '/': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Dashboard }))
    },
  },
  '/products': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Products }))
    },
  },
  '/products/new': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: CreateProduct }))
    },
  },
  '/products/:id': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Product }))
    },
  },
  '/coupons': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Coupons }))
    },
  },
  '/orders': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Orders }))
    },
  },
  '/orders/:id': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Order }))
    },
  },
  '/signin': {
    render: () => {
      return m(signIn)
    },
  },
})
