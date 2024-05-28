import m from 'mithril'
import API from '../api'
import { IOrder } from '../pkg/models'

interface IOrderView extends m.Component {
  fetch: (id: string) => void
  updateOrder: () => void
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
  updateOrder: (): void => {
    API.updateOrder(Order.order)
      .then((res) => res.data)
      .then((res) => {
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m(
      'main',
      {
        className: 'w-1/2',
      },
      Order.order ?
        [
          m('h3', Order.order.id),
          m('table', { className: 'w-full my-5 shadow rounded' }, [
            m('tr', { className: 'py-3' }, [
              m('th', 'Cost'),
              m('td', `${Order.order.currency} ${Order.order.totalCost}`),
            ]),
            m('tr', { className: 'py-3' }, [m('th', 'City'), m('td', Order.order.city)]),
            m('tr', { className: 'py-3' }, [m('th', 'Country'), m('td', Order.order.country)]),
            m('tr', { className: 'py-3' }, [
              m('th', 'Postal code'),
              m('td', Order.order.postalCode),
            ]),
            m('tr', { className: 'py-3' }, [m('th', 'Address'), m('td', Order.order.address)]),
          ]),
          m(
            'select',
            {
              className: 'py-5 px-2, w-full',
              onchange: (e: Event) => {
                Order.order.status = (e.target as HTMLSelectElement).value
              },
            },
            [
              m('option', 'placed'),
              m('option', 'paid'),
              m('option', 'being-prepered'),
              m('option', 'sent'),
            ],
          ),
          m('table', { className: 'w-full my-5 shadow rounded' }, [
            m(
              'thead',
              { className: 'bg-midnightGreen text-antiflashWhite rounded-t' },
              m('tr', [m('th', 'Name'), m('th', 'Price'), m('th', 'Quantity')]),
            ),
            m('tbody', [
              Object.entries(Order.order.basket).map(([_, value]) =>
                m('tr', { className: 'py-3 text-center' }, [
                  m('td', value.name),
                  m('td', `${value.currency} ${value.price}`),
                  m('td', value.quantity),
                ]),
              ),
            ]),
          ]),
          m('button', { onclick: () => Order.updateOrder() }, 'save'),
        ]
      : null,
    )
  },
  order: undefined,
}

export default Order
