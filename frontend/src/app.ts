import m from 'mithril'
import RequireAuth from './auth/RequireAuth'

// views
import signIn from './pages/signIn'
import Dashboard from './pages/dashboard'
import Products from './pages/products'
import Product from './pages/product'
import CreateProduct from './pages/createProduct'

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
  '/signin': {
    render: () => {
      return m(signIn)
    },
  },
})
