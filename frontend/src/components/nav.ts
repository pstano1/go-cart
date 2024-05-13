import m from 'mithril'

const Nav: m.Component = {
  view: () => {
    return m('nav', { class: 'flex p-4 shadow gap-4' }, [
      m('div', { class: 'flex-1' }),
      m('a', { href: '#!/products', class: 'cursor-pointer text-lg capitalize' }, 'products'),
      m('a', { href: '#!/coupons', class: 'cursor-pointer text-lg capitalize' }, 'coupons'),
      m('a', { href: '#!/orders', class: 'cursor-pointer text-lg capitalize' }, 'orders'),
    ])
  },
}

export default Nav
