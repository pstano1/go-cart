import m from 'mithril'
import API from '../api'
import { IProduct } from '../pkg/models'

interface IProductView extends m.Component {
  product: IProduct
  fetchProduct: (id: string) => void
}

const Product: IProductView = {
  oncreate: () => {
    const id = m.route.param('id')
    Product.fetchProduct(id)
  },
  fetchProduct: (id: string) => {
    API.getProducts(id)
      .then((res: any) => res.data)
      .then((res) => {
        Product.product = res[0]
        m.redraw()
      })
      .catch((err) => {
        // handle
      })
  },
  view: () => {
    return m('main', [m('h2', Product.product?.name)])
  },
  product: undefined,
}

export default Product
