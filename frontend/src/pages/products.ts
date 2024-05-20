import m from 'mithril'
import API from '../api'
import { IProduct } from '../pkg/models'

interface IProductView extends m.Component {
  products: IProduct[]
  categories: string[]
  saveCategory: () => void
  fetchProducts: () => void
  fetchCategories: () => void
}

const Products: IProductView = {
  oninit: () => {
    Products.fetchProducts()
    Products.fetchCategories()
  },
  fetchProducts: (): void => {
    API.getProducts()
      .then((res) => res.data)
      .then((res) => {
        Products.products = res
        m.redraw()
      })
      .catch((err) => {
        // handle error
      })
  },
  fetchCategories: (): void => {
    API.getCategories()
      .then((res) => res.data)
      .then((res) => {
        Products.categories = res
        m.redraw()
      })
      .catch((err) => {
        // handle error
      })
  },
  saveCategory: (): void => {
    const categoryInput = document.getElementById('category-textfield') as HTMLInputElement
    const name: string = categoryInput.value.replace(/[^a-zA-Z0-9\s]/g, '').replace(/\s+/g, '-')

    API.createCategory({ name: name })
      .then((res) => res.data)
      .then((res) => {
        // pass
      })
      .then(() => (categoryInput.value = ''))
      .catch((err) => {
        // handle error
      })
  },
  view: () => {
    return m('main', [
      m('nav', { className: 'flex p-4 shadow gap-4 mb-2 rounded sticky top-2 bg-white z-999' }, [
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
            { className: 'shadow rounded overflow-hidden' },
            m('img', {
              className: 'object-cover w-full max-h-72',
              src: './images/gfx/product.jpg',
            }),
            m('div', { className: 'p-5' }, [
              m('div', { className: 'flex space-between items-center' }, [
                m('h4', { className: 'my-2 text-xl text-bolder' }, product.names['PL']),
                m('span', { className: 'flex-1' }),
                m(
                  'div',
                  {
                    className:
                      'px-3 py-1 rounded-3xl bg-midnightGreen w-fit h-fit text-antiflashWhite',
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
              m('p', { className: 'text-justify my-2' }, product.descriptions['PL']),
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
          m('tbody', [
            Products.categories.map((category) =>
              m('tr', [m('td', { className: 'py-3 text-center w-1/2' }, category), m('td')]),
            ),
            m('tr', [
              m(
                'td',
                m('input', {
                  type: 'text',
                  id: 'category-textfield',
                  className: 'shadow block my-2 text-lg mx-2 py-3 px-2 rounded w-full',
                }),
              ),
              m(
                'button',
                {
                  onclick: (event: Event) => {
                    event.preventDefault()
                    Products.saveCategory()
                  },
                  className: 'block my-auto mx-auto',
                },
                'save',
              ),
            ]),
          ]),
        ]),
      ),
    ])
  },
  products: [],
  categories: [],
}

export default Products
