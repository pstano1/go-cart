import m from 'mithril'
import RequireAuth from './auth/RequireAuth'

// views
import signIn from './pages/signIn'
import Dashboard from './pages/dashboard'
import Products from './pages/products'

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
  '/signin': {
    render: () => {
      return m(signIn)
    },
  },
})
