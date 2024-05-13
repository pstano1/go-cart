import m from 'mithril'
import Nav from './nav'

const Layout = {
  view: (vnode: m.Vnode<any, any>) => {
    return m('div', [m(Nav), m('main', { className: 'p-5' }, m(vnode.attrs.contentComponent))])
  },
}

export default Layout
