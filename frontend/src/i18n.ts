import i18n from 'i18next'

import commonPL from '../bin/dictionaries/common/pl.json'
import commonEN from '../bin/dictionaries/common/en.json'
import productsPL from '../bin/dictionaries/products/pl.json'
import productsEN from '../bin/dictionaries/products/en.json'
import couponsPL from '../bin/dictionaries/coupons/pl.json'
import couponsEN from '../bin/dictionaries/coupons/en.json'
import ordersPL from '../bin/dictionaries/orders/pl.json'
import ordersEN from '../bin/dictionaries/orders/en.json'

const language: string = localStorage.getItem('language')
i18n.init({
  fallbackLng: 'en',
  lng: language,
  ns: ['Common'],
  defaultNS: 'Common',
  supportedLngs: ['en', 'pl'],
  resources: {
    pl: {
      Common: { ...commonPL },
      Products: { ...productsPL },
      Coupons: { ...couponsPL },
      Orders: { ...ordersPL },
    },
    en: {
      Common: { ...commonEN },
      Products: { ...productsEN },
      Coupons: { ...couponsEN },
      Orders: { ...ordersEN },
    },
  },
  interpolation: {
    escapeValue: false,
  },
})
