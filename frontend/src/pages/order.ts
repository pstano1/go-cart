import m from 'mithril'
import API from '../api'
import { IBasketEntry, IOrder } from '../pkg/models'

interface IOrderView extends m.Component {
  fetch: (id: string) => void
  order: IOrder
}

const Order: IOrderView = {
  oncreate: (): void => {
    const id = m.route.param('id')
    Order.fetch(id)
  },
  fetch: (id: string): void => {
    API.getOrders(id)
      .then((res) => res.data)
      .then((res) => {
        Order.order = res[0]
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m(
      'main',
      Order.order ?
        [
          m('h3', Order.order.id),
          m('table', [
            m('tr', [m('th', 'Cost'), m('td', `${Order.order.currency} ${Order.order.totalCost}`)]),
            m('tr', [m('th', 'City'), m('td', Order.order.city)]),
            m('tr', [m('th', 'Country'), m('td', Order.order.country)]),
            m('tr', [m('th', 'Postal code'), m('td', Order.order.postalCode)]),
            m('tr', [m('th', 'Address'), m('td', Order.order.address)]),
          ]),
          m('table', [
            m('table', [
              m('thead', m('tr', [m('th', 'Name'), m('th', 'Price'), m('th', 'Quantity')])),
              m('tbody', [
                Object.entries(Order.order.basket).map(([_, value]) =>
                  m('tr', [
                    m('td', value.name),
                    m('td', `${value.currency} ${value.price}`),
                    m('td', value.quantity),
                  ]),
                ),
              ]),
            ]),
          ]),
        ]
      : null,
    )
  },
  order: undefined,
}

export default Order
