import m from 'mithril'
import UserIcon from '../../bin/images/icons/user.svg'

interface INav extends m.Component {
  isUserMenuOpen: boolean
  signOut: () => void
}

const Nav: INav = {
  signOut: (): void => {
    localStorage.setItem('sessionToken', null)
    Nav.isUserMenuOpen = false
    m.route.set('/signin')
  },
  view: () => {
    return m('nav', { class: 'flex p-4 shadow gap-4 items-center' }, [
      m('div', { class: 'flex-1' }),
      m('a', { href: '#!/products', class: 'cursor-pointer text-lg capitalize' }, 'products'),
      m('a', { href: '#!/coupons', class: 'cursor-pointer text-lg capitalize' }, 'coupons'),
      m('a', { href: '#!/orders', class: 'cursor-pointer text-lg capitalize' }, 'orders'),
      m('div', { className: 'relative' }, [
        m(
          'div',
          {
            className: 'rounded-full bg-azure p-2 cursor-pointer',
            onclick: () => {
              Nav.isUserMenuOpen = !Nav.isUserMenuOpen
            },
          },
          m('img', { src: UserIcon }),
        ),
        Nav.isUserMenuOpen &&
          m(
            'div',
            { className: 'absolute bg-white rounded-lg shadow py-3 right-1 w-24' },
            m('ul', [
              m(
                'li',
                { className: 'cursor-pointer py-2 px-4 hover:bg-azure', onclick: Nav.signOut },
                'Sign out',
              ),
            ]),
          ),
      ]),
    ])
  },
  isUserMenuOpen: false,
}

export default Nav
