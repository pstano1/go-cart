import m from 'mithril'
import API from '../api'
import { ICoupon } from '../pkg/models'

interface ICouponsView extends m.Component {
  fetch: () => void
  coupons: ICoupon[]
}

const Coupons: ICouponsView = {
  oninit: () => {
    Coupons.fetch()
  },
  fetch: (): void => {
    API.getCoupons()
      .then((res) => res.data)
      .then((res) => {
        Coupons.coupons = res
      })
      .catch((err) => {})
  },
  view: () => {
    return m('main', [
      m('table', { className: 'border-collapse shadow w-full' }, [
        m('thead', [
          m('tr', [m('th', 'Code'), m('th', 'Amount'), m('th', 'Unit'), m('th', 'Actions')]),
        ]),
        m('tbody', [
          Coupons.coupons?.map((coupon) =>
            m('tr', [m('td', coupon.promoCode), m('td', coupon.amount), m('td', coupon.unit)]),
          ),
        ]),
      ]),
    ])
  },
  coupons: [],
}

export default Coupons
