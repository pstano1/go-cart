import m from 'mithril'

const Nav: m.Component = {
  view: () => {
    return m('nav', { class: 'flex p-4 shadow' }, [
      m('div', { class: 'flex-1' }),
      m('a', { href: '#!/products', class: 'cursor-pointer text-lg' }, 'products'),
    ])
  },
}

export default Nav
