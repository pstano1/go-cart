import m from 'mithril'
import API from '../api'
import { t } from 'i18next'
import { IOrder } from '../pkg/models'
import InfoIcon from '../../bin/images/icons/info.svg'

interface IOrdersView extends m.Component {
  orders: IOrder[]
  fetch: () => void
}

const Orders: IOrdersView = {
  oncreate: (): void => {
    Orders.fetch()
  },
  fetch: (): void => {
    API.getOrders()
      .then((res) => res.data)
      .then((res) => {
        Orders.orders = res
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m(
      'main',
      m('table', { className: 'border-collapse shadow w-full' }, [
        m('thead', [
          m('tr', [
            m('th', { className: 'py-2' }, 'Id'),
            m('th', { className: 'py-2' }, t('Orders:status')),
            m('th', { className: 'py-2' }, t('Orders:actions')),
          ]),
        ]),
        m(
          'tbody',
          Orders.orders?.map((order) =>
            m('tr', { className: 'cursor-pointer' }, [
              m('td', { className: 'p-2' }, order.id),
              m('td', { className: 'p-2' }, order.status),
              m(
                'td',
                { className: 'p-2' },
                m(
                  m.route.Link,
                  {
                    href: `/orders/${order.id}`,
                  },
                  m('img', { src: InfoIcon }),
                ),
              ),
            ]),
          ),
        ),
      ]),
    )
  },
  orders: [],
}

export default Orders
