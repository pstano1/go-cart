import m from 'mithril'
import API from '../api'
import { t } from 'i18next'
import { ICategory, IProduct } from '../pkg/models'

import AddIcon from '../../bin/images/icons/plus.svg'
import SaveIcon from '../../bin/images/icons/save.svg'
import EditIcon from '../../bin/images/icons/edit-2.svg'
import DeleteIcon from '../../bin/images/icons/trash-2.svg'

interface IProductView extends m.Component {
  products: IProduct[]
  categories: ICategory[]
  categoryInEdition: string
  saveCategory: () => void
  fetchProducts: () => void
  fetchCategories: () => void
  updateCategory: (category: ICategory) => void
  deleteCategory: (id: string) => void
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
  updateCategory: (category: ICategory): void => {
    API.updateCategory(category)
      .then((res) => {})
      .then(() => {
        m.redraw()
      })
      .catch((err) => {})
  },
  deleteCategory: (id: string): void => {
    API.deleteCategory(id)
      .then((res) => {})
      .then(() => {
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m('main', [
      m('nav', { className: 'flex p-4 shadow gap-4 mb-2 rounded sticky top-2 bg-white z-49' }, [
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
          t('Products:subnav.products'),
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
          t('Products:subnav.categories'),
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
                t('Products:more'),
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
              m('th', { className: 'py-2' }, t('Products:name')),
              m('th', { className: 'py-2' }, t('Products:actions')),
            ]),
          ),
          m('tbody', [
            Products.categories.map((category) =>
              m('tr', [
                m(
                  'td',
                  { className: 'py-3 text-center w-1/2' },
                  Products.categoryInEdition === category.id ?
                    m('input', {
                      id: 'new-category-name',
                      value: category.name,
                      className: 'shadow block my-2 text-lg mx-2 py-3 px-2 rounded w-full',
                    })
                  : category.name,
                ),
                m(
                  'td',
                  { className: 'flex items-center justify-center gap-5' },
                  Products.categoryInEdition === category.id ?
                    m(
                      'button',
                      {
                        onclick: (event: Event) => {
                          event.preventDefault()
                          Products.updateCategory({
                            name: (document.getElementById('new-category-name') as HTMLInputElement)
                              .value,
                            id: category.id,
                          })
                        },
                        className: 'block my-auto mx-auto',
                      },
                      m('img', { src: SaveIcon }),
                    )
                  : [
                      m(
                        'button',
                        {
                          onclick: (event: Event) => {
                            event.preventDefault()
                            Products.categoryInEdition = category.id
                          },
                        },
                        m('img', { src: EditIcon }),
                      ),
                      m(
                        'button',
                        {
                          onclick: (event: Event) => {
                            event.preventDefault()
                            Products.deleteCategory(category.id)
                          },
                        },
                        m('img', { src: DeleteIcon }),
                      ),
                    ],
                ),
              ]),
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
                'td',
                m(
                  'button',
                  {
                    onclick: (event: Event) => {
                      event.preventDefault()
                      Products.saveCategory()
                    },
                    className: 'block my-auto mx-auto',
                  },
                  m('img', { src: AddIcon }),
                ),
              ),
            ]),
          ]),
        ]),
      ),
    ])
  },
  products: [],
  categories: [],
  categoryInEdition: '',
}

export default Products
