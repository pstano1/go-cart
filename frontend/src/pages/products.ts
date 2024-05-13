import m from 'mithril'
import API from '../api'
import { IProduct } from '../pkg/models'

interface IProductView extends m.Component {
  products: IProduct[]
  fetchProducts: () => void
}

const Products: IProductView = {
  oninit: () => {
    Products.fetchProducts()
  },
  fetchProducts: () => {
    API.getProducts()
      .then((res: any) => res.data)
      .then((res) => {
        Products.products = res
        m.redraw()
      })
      .catch((err) => {
        // handle error
      })
  },
  view: () => {
    return m('main', [
      m('nav', { className: 'flex p-4 shadow gap-4 mb-2 rounded sticky top-2 bg-white' }, [
        m(
          'a',
          {
            className: 'cursor-pointer text-lg capitalize',
            href: '#!/products/#products',
            onclick: (event: Event): void => {
              const productsSection: HTMLElement = document.getElementById('products')

              if (productsSection) {
                event.preventDefault()
                productsSection.scrollIntoView({ behavior: 'smooth' })
                return
              }
            },
          },
          'products',
        ),
        m(
          'a',
          {
            className: 'cursor-pointer text-lg capitalize',
            href: '#!/products/#categories',
            onclick: (event: Event): void => {
              const categoriesSection: HTMLElement = document.getElementById('categories')

              if (categoriesSection) {
                event.preventDefault()
                categoriesSection.scrollIntoView({ behavior: 'smooth' })
                return
              }
            },
          },
          'categories',
        ),
      ]),
      m(
        'section',
        { id: 'products', className: 'grid md:grid-cols-3 gap-3 grid-cols-1' },
        Products.products.map((product) =>
          m(
            'div',
            { className: 'shadow rounded overflow-hidden shadow' },
            m('img', {
              className: 'object-cover w-full max-h-72',
              src: './images/gfx/product.jpg',
            }),
            m('div', { className: 'p-5' }, [
              m('div', { className: 'flex space-between' }, [
                m('h4', { className: 'my-2 text-xl text-bolder' }, product.name),
                m('span', { className: 'flex-1' }),
                m(
                  'div',
                  {
                    className:
                      'px-3 py-1 rounded-3xl bg-midnightGreen w-fit h-fit text-antiflashWhite relative top-1.5',
                  },
                  `PLN ${product.prices['PLN']}`,
                ),
              ]),
              m(
                'div',
                { className: 'overflow-y-scroll flex gap-2 my-2' },
                product.categories.map((category) =>
                  m(
                    'p',
                    {
                      className: 'px-3 py-1 rounded-3xl bg-midnightGreen w-fit text-antiflashWhite',
                    },
                    category,
                  ),
                ),
              ),
              m('p', { className: 'justify my-2' }, product.descriptions['PL']),
              m(
                m.route.Link,
                {
                  className:
                    'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto block w-fit my-2',
                  href: `/products/${product.id}`,
                },
                'more',
              ),
            ]),
          ),
        ),
      ),
      m(
        m.route.Link,
        {
          className: 'py-5 shadow my-2 w-full text-center text-7xl h-fit cursor-pointer block',
          href: '/products/new',
        },
        '+',
      ),
      m(
        'section',
        { id: 'categories', className: 'my-2' },
        m('table', { className: 'border-collapse shadow w-full' }, [
          m(
            'thead',
            m('tr', [
              m('th', { className: 'py-2' }, 'Name'),
              m('th', { className: 'py-2' }, 'Actions'),
            ]),
          ),
          m('tbody', [m('tr')]),
        ]),
      ),
    ])
  },
  products: [],
}

export default Products
