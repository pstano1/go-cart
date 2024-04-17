import m from 'mithril'
import signIn from './pages/signIn'

const Layout = {
  view: (vnode: m.Vnode<any, any>) => {
    return m('div', [m(vnode.attrs.contentComponent)])
  },
}

const Main = {
  view: () => {
    return m('div', { class: 'text-3xl' }, 'go-cart', m(m.route.Link, { href: '/signin', options: {replace: true} }, 'sign in'))
  }
}

m.route(document.body, '/signin', {
  '/': {
    render: () => {
      return m(Layout, {
        contentComponent: Main,
      })
    },
  },
  '/signin': {
    render: () => {
      return m(signIn)
    }
  },
})
