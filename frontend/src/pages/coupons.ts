import m from 'mithril'
import API from '../api'
import { t } from 'i18next'
import { ICoupon } from '../pkg/models'
import { CouponCreate, CouponUpdate } from '../pkg/requests'

import AddIcon from '../../bin/images/icons/plus.svg'
import DisableIcon from '../../bin/images/icons/slash.svg'
import EnableIcon from '../../bin/images/icons/target.svg'
import DeleteIcon from '../../bin/images/icons/trash-2.svg'

interface ICouponsView extends m.Component {
  fetch: () => void
  coupons: ICoupon[]
  saveCoupon: () => void
  updateCoupon: (req: CouponUpdate) => void
  deleteCoupon: (id: string) => void
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
        m.redraw()
      })
      .catch((err) => {})
  },
  saveCoupon: (): void => {
    const couponForm = document.getElementById('coupon-create-form') as HTMLElement
    const inputs = couponForm.querySelectorAll('input')
    let coupon: CouponCreate = {
      promoCode: '',
      amount: 0,
      unit: '',
    }

    inputs.forEach((input) => {
      if (input.id) {
        switch (input.id) {
          case 'promoCode':
            coupon.promoCode = input.value
            break
          case 'amount':
            coupon.amount = Number(input.value)
            break
          case 'unit':
            coupon.unit = input.value
            break
          default:
            break
        }
      }
    })

    API.createCoupon(coupon)
      .then((res) => res.data)
      .then((res) => {
        Coupons.fetch()
        m.redraw()
      })
      .catch((err) => {})
  },
  updateCoupon: (req: CouponUpdate) => {
    API.updateCoupon(req)
      .then(() => {
        Coupons.fetch()
        m.redraw()
      })
      .catch((err) => {})
  },
  deleteCoupon: (id: string): void => {
    API.deleteCoupon(id)
      .then(() => {
        Coupons.fetch()
        m.redraw()
      })
      .catch((err) => {})
  },
  view: () => {
    return m('main', [
      m('table', { className: 'border-collapse shadow w-full' }, [
        m('thead', [
          m('tr', [
            m('th', { className: 'py-2' }, t('Coupons:code')),
            m('th', { className: 'py-2' }, t('Coupons:amount')),
            m('th', { className: 'py-2' }, t('Coupons:unit')),
            m('th', { className: 'py-2' }, t('Coupons:actions')),
          ]),
        ]),
        m('tbody', [
          Coupons.coupons?.map((coupon) =>
            m('tr', [
              m('td', coupon.promoCode),
              m('td', coupon.amount),
              m('td', coupon.unit),
              m('td', [
                coupon.isActive ?
                  m(
                    'button',
                    {
                      onclick: (event: Event) => {
                        event.preventDefault()
                        let req: CouponUpdate = coupon
                        req.isActive = false
                        Coupons.updateCoupon(req)
                      },
                      className: 'block my-auto mx-auto',
                    },
                    m('img', { src: DisableIcon }),
                  )
                : m(
                    'button',
                    {
                      onclick: (event: Event) => {
                        event.preventDefault()
                        let req: CouponUpdate = coupon
                        req.isActive = false
                        Coupons.updateCoupon(req)
                      },
                      className: 'block my-auto mx-auto',
                    },
                    m('img', { src: EnableIcon }),
                  ),
                m(
                  'button',
                  {
                    onclick: (event: Event) => {
                      event.preventDefault()
                      Coupons.deleteCoupon(coupon.id)
                    },
                    className: 'block my-auto mx-auto',
                  },
                  m('img', { src: DeleteIcon }),
                ),
              ]),
            ]),
          ),
          m(
            'tr',
            {
              id: 'coupon-create-form',
            },
            [
              m(
                'td',
                m('input', {
                  type: 'text',
                  id: 'promoCode',
                  className: 'shadow block my-2 text-lg mx-2 py-3 px-2 rounded w-full',
                }),
              ),
              m(
                'td',
                m('input', {
                  type: 'number',
                  id: 'amount',
                  className: 'shadow block my-2 text-lg mx-2 py-3 px-2 rounded w-full',
                }),
              ),
              m(
                'td',
                m('input', {
                  type: 'text',
                  id: 'unit',
                  className: 'shadow block my-2 text-lg mx-2 py-3 px-2 rounded w-full',
                }),
              ),
              m(
                'td',
                m(
                  'button',
                  {
                    onclick: (event: Event) => {
                      event.preventDefault()
                      Coupons.saveCoupon()
                    },
                    className: 'block my-auto mx-auto',
                  },
                  m('img', { src: AddIcon }),
                ),
              ),
            ],
          ),
        ]),
      ]),
    ])
  },
  coupons: [],
}

export default Coupons
