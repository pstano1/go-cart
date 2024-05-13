import m from 'mithril'
import API from '../api'

interface IProductView extends m.Component {
  products: any[]
  fetchProducts: () => void
}

const Products: IProductView = {
  oncreate: () => {
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
        // handle
      })
  },
  view: () => {
    return m('main', [
      m(
        'section',
        { className: 'grid grid-cols-3 gap-3' },
        Products.products.map((product) =>
          m('div', { className: 'shadow p-3 rounded' }, [m('h4', product.name)]),
        ),
      ),
    ])
  },
  products: [],
}

export default Products
