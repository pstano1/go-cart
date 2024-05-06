import m from 'mithril'
import RequireAuth from './auth/RequireAuth'

// views
import signIn from './pages/signIn'
import Dashboard from './pages/dashboard'

import Nav from './components/nav'

const Layout = {
  view: (vnode: m.Vnode<any, any>) => {
    return m('div', [m(Nav), m(vnode.attrs.contentComponent)])
  },
}

m.route(document.body, '/signin', {
  '/': {
    render: () => {
      return m(RequireAuth, m(Layout, { contentComponent: Dashboard }))
    },
  },
  '/signin': {
    render: () => {
      return m(signIn)
    },
  },
})
