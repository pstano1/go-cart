import m from 'mithril'

const Layout = {
  view: (vnode: m.Vnode<any, any>) => {
    return m('div', [m(vnode.attrs.contentComponent)])
  },
}

m.route(document.body, '/', {
  '/': {
    render: () => {
      return m(Layout, {
        contentComponent: {
          view: () => {
            return m('div', { class: 'text-3xl' }, 'go-cart')
          },
        },
      })
    },
  },
})
