import m from 'mithril'
import API from '../api'
import { IProduct } from '../pkg/models'
import { ProductUpdate } from '../pkg/requests'

interface IProductView extends m.Component {
  product: Partial<IProduct>
  fetchProduct: (id: string) => void
  languages: string[]
  currencies: string[]
  addToLanguages: (languague: string) => void
  addToCurrencies: (currency: string) => void
  handleSubmit: (event: Event) => void
}

const Product: IProductView = {
  oncreate: () => {
    Product.currencies = []
    Product.languages = []
    const id = m.route.param('id')
    Product.fetchProduct(id)
  },
  handleSubmit: (event: Event): void => {
    event.preventDefault()

    const formData = new FormData(event.target as HTMLFormElement)
    let product: ProductUpdate = {
      id: Product.product?.id,
      names: {},
      categories: ['test', 'test1', 'test2', 'broad-category', 'test3'],
      descriptions: {},
      prices: {},
    }

    for (let [key, value] of formData.entries()) {
      if (key === 'categories') {
        // pass
      }
      if (key.includes('description')) {
        key = key.replace('description-', '')
        product.descriptions[key] = value as string
      }
      if (key.includes('name')) {
        key = key.replace('name-', '')
        product.names[key] = value as string
      }
      if (key.length === 3) {
        product.prices[key] = Number(value)
      }
    }

    API.updateProduct(product)
      .then((res) => res.data)
      .then((res) => {
        // TODO: display toast
        Product.product = product
      })
      .catch((err) => {
        // handle error
      })
  },
  fetchProduct: (id: string) => {
    API.getProducts(id)
      .then((res: any) => res.data)
      .then((res) => {
        Product.product = res[0]
        Object.entries(Product.product?.descriptions || {}).map(([key, _]) => {
          Product.languages.push(key)
        })
        Object.entries(Product.product?.prices || {}).map(([key, _]) => {
          Product.currencies.push(key)
        })
        m.redraw()
      })
      .catch((err) => {
        // handle
      })
  },
  addToLanguages: (languague: string) => {
    if (Product.languages.includes(languague.toUpperCase())) {
      return
    }
    Product.languages.push(languague.toUpperCase())
    m.redraw()
  },
  addToCurrencies: (currency: string) => {
    if (Product.currencies.includes(currency.toUpperCase())) {
      return
    }
    Product.currencies.push(currency.toUpperCase())
    m.redraw()
  },
  view: () => {
    return m('main', [
      m('section', { className: 'w-1/2' }, [
        m('h2', { className: 'text-2xl' }, 'Product overview'),
        m('div', { className: 'block w-1/2 h-fit' }, [
          m('label', { className: 'text-lg text-bolder block' }, 'Languages'),
          m(
            'div',
            { className: 'flex gap-1' },
            Product.languages.map((languague) =>
              m(
                'div',
                { className: 'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded w-fit' },
                languague,
              ),
            ),
          ),
          m('input', {
            id: 'languageField',
            type: 'text',
            minlength: 2,
            maxlength: 2,
            className: 'shadow inline-block my-4 text-lg w-full py-3 px-2 rounded w-fit',
          }),
          m(
            'button',
            {
              onclick: (event: Event) => {
                event.preventDefault()

                const languague = (document.getElementById('languageField') as HTMLInputElement)
                  .value
                Product.addToLanguages(languague)
              },
              className:
                'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto inline-block w-fit mx-2 h-full',
            },
            'add',
          ),
        ]),
        m('form', { className: 'block h-fit', name: 'product', onsubmit: Product.handleSubmit }, [
          m(
            'div',
            Product.languages?.map((language) => [
              m('label', { className: 'text-lg' }, language),
              m('input', {
                name: 'name-' + language,
                className: 'shadow inline-block my-4 text-lg w-full py-3 px-2 rounded w-fit',
                value: Product.product?.names[language],
              }),
            ]),
          ),
          m('label', { className: 'text-lg text-bolder' }, 'Description'),
          m(
            'div',
            Product.languages?.map((language) => [
              m('label', { className: 'text-lg' }, language),
              m('textarea', {
                maxlength: 250,
                name: 'description-' + language,
                className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded resize-none',
                value: Product.product?.descriptions[language],
              }),
            ]),
          ),
          m('label', { className: 'text-lg text-bolder' }, 'Price'),
          m('div', { className: 'block w-1/2 h-fit' }, [
            m('input', {
              id: 'currencyField',
              type: 'text',
              minlength: 3,
              maxlength: 3,
              className: 'shadow inline-block my-4 text-lg w-full py-3 px-2 rounded w-fit',
            }),
            m(
              'button',
              {
                onclick: (event: Event) => {
                  event.preventDefault()

                  const currency = (document.getElementById('currencyField') as HTMLInputElement)
                    .value
                  Product.addToCurrencies(currency)
                },
                className:
                  'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto inline-block w-fit mx-2 h-full',
              },
              'add',
            ),
          ]),
          m(
            'div',
            Product.currencies?.map((currency) => [
              m('label', { className: 'text-lg' }, currency),
              m('input', {
                type: 'number',
                name: currency,
                className: 'shadow inline-block my-4 text-lg w-full py-3 px-2 rounded w-fit',
                value: Product.product?.prices[currency],
              }),
            ]),
          ),
          m(
            'button',
            {
              type: 'submit',
              className:
                'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto block w-fit my-2',
            },
            'update',
          ),
        ]),
        m(
          'button',
          {
            onclick: (event: Event) => {
              event.preventDefault()

              API.deleteProduct(Product.product?.id)
                .then((res) => res.data)
                .then(() => {
                  // TODO: display success toast
                  window.history.go(-1)
                })
                .catch((err) => {
                  // handle error
                })
            },
          },
          'delete',
        ),
      ]),
    ])
  },
  product: undefined,
  languages: [],
  currencies: [],
}

export default Product
