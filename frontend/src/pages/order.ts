import m from 'mithril'
import API from '../api'
import { t } from 'i18next'
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
              m('th', t('Orders:cost')),
              m('td', `${Order.order.currency} ${Order.order.totalCost}`),
            ]),
            m('tr', { className: 'py-3' }, [m('th', t('Orders:city')), m('td', Order.order.city)]),
            m('tr', { className: 'py-3' }, [
              m('th', t('Orders:country')),
              m('td', Order.order.country),
            ]),
            m('tr', { className: 'py-3' }, [
              m('th', 'Postal code'),
              m('td', Order.order.postalCode),
            ]),
            m('tr', { className: 'py-3' }, [
              m('th', t('Orders:address')),
              m('td', Order.order.address),
            ]),
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
              m('option', { value: 'placed' }, t('Orders:statuses.placed')),
              m('option', { value: 'cancelled' }, t('Orders:statuses.cancelled')),
              m('option', { value: 'paid' }, t('Orders:statuses.paid')),
              m('option', { value: 'being-prepared' }, t('Orders:statuses.being-prepared')),
              m('option', { value: 'sent' }, t('Orders:statuses.sent')),
            ],
          ),
          m('table', { className: 'w-full my-5 shadow rounded' }, [
            m(
              'thead',
              { className: 'bg-midnightGreen text-antiflashWhite rounded-t' },
              m('tr', [
                m('th', t('Orders:name')),
                m('th', t('Orders:price')),
                m('th', t('Orders:quantity')),
              ]),
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
          m('button', { onclick: () => Order.updateOrder() }, t('Orders:save')),
        ]
      : null,
    )
  },
  order: undefined,
}

export default Order
