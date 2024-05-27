export interface IProduct {
  id: string
  names: {
    [key: string]: string
  }
  categories: string[]
  descriptions: {
    [key: string]: string
  }
  prices: {
    [key: string]: number
  }
  priceHistory: {
    [key: string]: number
  }
}

export interface ICategory {
  id: string
  name: string
}

export interface ICoupon {
  id: string
  promoCode: string
  amount: number
  unit: string
  isActive: boolean
}

export interface IBasketEntry {
  name: string
  quantity: number
  currency: string
  price: string
}

export interface IOrder {
  id: string
  customerId: string
  totalCost: number
  currency: string
  country: string
  city: string
  postalCode: string
  address: string
  status: string
  basket: {
    [key: string]: IBasketEntry
  }
  taxId?: string
}
