import m from 'mithril'
import API from '../api'
import { ProductCreate } from '../pkg/requests'

interface ICreateProductView extends m.Component {
  languages: string[]
  currencies: string[]
  addToLanguages: (languague: string) => void
  addToCurrencies: (currency: string) => void
  handleSubmit: (event: Event) => void
}

const CreateProduct: ICreateProductView = {
  handleSubmit: (event: Event): void => {
    event.preventDefault()

    const formData = new FormData(event.target as HTMLFormElement)
    let newProduct: ProductCreate = {
      name: '',
      categories: ['test', 'test1', 'test2', 'broad-category', 'test3'],
      descriptions: {},
      prices: {},
    }

    for (let [key, value] of formData.entries()) {
      if (key === 'categories') {
        // pass
      }
      if (key.length === 2) {
        newProduct.descriptions[key] = value as string
      }
      if (key.length === 3) {
        newProduct.prices[key] = Number(value)
      }
      if (key === 'name') {
        newProduct.name = value as string
      }
    }

    API.createProduct(newProduct)
      .then((res) => res.data)
      .then((res) => {
        // pass
      })
      .catch((err) => {
        // handle error
      })
  },
  addToLanguages: (languague: string) => {
    if (CreateProduct.languages.includes(languague.toUpperCase())) {
      return
    }
    CreateProduct.languages.push(languague.toUpperCase())
    m.redraw()
  },
  addToCurrencies: (currency: string) => {
    if (CreateProduct.currencies.includes(currency.toUpperCase())) {
      return
    }
    CreateProduct.currencies.push(currency.toUpperCase())
    m.redraw()
  },
  view: () => {
    return m('main', [
      m('h2', { className: 'text-2xl' }, 'Add a new product'),
      m('div', { className: 'block w-1/2 h-fit' }, [
        m('label', { className: 'text-lg text-bolder block' }, 'Languages'),
        m(
          'div',
          { className: 'flex gap-1' },
          CreateProduct.languages.map((languague) =>
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
            onclick: () => {
              const languague = (document.getElementById('languageField') as HTMLInputElement).value
              CreateProduct.addToLanguages(languague)
            },
            className:
              'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto inline-block w-fit mx-2 h-full',
          },
          'add',
        ),
      ]),
      m(
        'form',
        {
          name: 'new-product',
          onsubmit: CreateProduct.handleSubmit,
          className: 'block w-1/2',
        },
        [
          m('label', { className: 'text-lg text-bolder' }, 'Name'),
          m('input', {
            type: 'text',
            name: 'name',
            className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded',
          }),
          m('label', { className: 'text-lg text-bolder' }, 'Categories'),
          m('select', { className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded' }),
          m('label', { className: 'text-lg text-bolder' }, 'Description'),
          m(
            'div',
            CreateProduct.languages.map((language) => [
              m('label', { className: 'text-lg' }, language),
              m('textarea', {
                maxlength: 250,
                name: language,
                className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded resize-none',
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
                  CreateProduct.addToCurrencies(currency)
                },
                className:
                  'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto inline-block w-fit mx-2 h-full',
              },
              'add',
            ),
          ]),
          m(
            'div',
            { className: 'block' },
            CreateProduct.currencies.map((currency) =>
              m(
                'div',
                { className: 'flex items-center' },
                m('label', { className: 'text-lg' }, currency),
                m('input', {
                  type: 'number',
                  name: currency,
                  className: 'shadow flex-1 my-4 text-lg w-full py-3 px-2 rounded mx-2 w-fit',
                }),
              ),
            ),
          ),
          m(
            'button',
            {
              type: 'submit',
              className:
                'bg-midnightGreen text-antiflashWhite px-3 py-1 rounded capitalize cursor-pointer mx-auto block w-fit my-2',
            },
            'save',
          ),
        ],
      ),
    ])
  },
  languages: [],
  currencies: [],
}

export default CreateProduct
