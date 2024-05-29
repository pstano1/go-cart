import m from 'mithril'
import API from '../api'
import { t } from 'i18next'
import { ProductCreate } from '../pkg/requests'
import { ICategory } from '../pkg/models'

interface ICreateProductView extends m.Component {
  languages: string[]
  currencies: string[]
  categories: ICategory[]
  addToLanguages: (languague: string) => void
  addToCurrencies: (currency: string) => void
  handleSubmit: (event: Event) => void
  fetchCategories: () => void
}

const CreateProduct: ICreateProductView = {
  oncreate: (): void => {
    CreateProduct.fetchCategories()
  },
  handleSubmit: (event: Event): void => {
    event.preventDefault()

    const formData = new FormData(event.target as HTMLFormElement)
    let newProduct: ProductCreate = {
      names: {},
      categories: [],
      descriptions: {},
      prices: {},
    }
    for (let [key, value] of formData.entries()) {
      if (key === 'categories') {
        const categorySelect = document.getElementById('category-select') as HTMLSelectElement
        const selectedOptions = Array.from(categorySelect.selectedOptions)
        newProduct.categories = selectedOptions.map((option) => option.value)
      }
      if (key.includes('description')) {
        key = key.replace('description-', '')
        newProduct.descriptions[key] = value as string
      }
      if (key.includes('name')) {
        key = key.replace('name-', '')
        newProduct.names[key] = value as string
      }
      if (key.length === 3) {
        newProduct.prices[key] = Number(value)
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
  fetchCategories: (): void => {
    API.getCategories()
      .then((res) => res.data)
      .then((res) => {
        CreateProduct.categories = res
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m('main', [
      m('h2', { className: 'text-2xl' }, t('Products:createPageTitle')),
      m('div', { className: 'block w-1/2 h-fit' }, [
        m('label', { className: 'text-lg text-bolder block' }, t('Products:languages')),
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
          t('Products:add'),
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
          m('label', { className: 'text-lg text-bolder' }, t('Products:name')),
          m(
            'div',
            CreateProduct.languages?.map((language) => [
              m('label', { className: 'text-lg' }, language),
              m('input', {
                name: 'name-' + language,
                className: 'shadow inline-block my-4 text-lg w-full py-3 px-2 rounded w-fit',
              }),
            ]),
          ),
          m('label', { className: 'text-lg text-bolder' }, t('Products:categories')),
          m(
            'select',
            {
              id: 'category-select',
              multiple: true,
              name: 'categories',
              className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded',
            },
            [
              CreateProduct.categories?.map((category) =>
                m('option', { value: category.name }, category.name),
              ),
            ],
          ),
          m('label', { className: 'text-lg text-bolder' }, t('Products:description')),
          m(
            'div',
            CreateProduct.languages.map((language) => [
              m('label', { className: 'text-lg' }, language),
              m('textarea', {
                maxlength: 250,
                name: 'description-' + language,
                className: 'shadow block my-4 text-lg w-full py-3 px-2 rounded resize-none',
              }),
            ]),
          ),
          m('label', { className: 'text-lg text-bolder' }, t('Products:price')),
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
              t('Products:add'),
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
            t('Products:save'),
          ),
        ],
      ),
    ])
  },
  languages: [],
  currencies: [],
  categories: [],
}

export default CreateProduct
