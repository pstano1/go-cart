export interface Credentials {
  customerId?: string
  username: string
  password: string
}

export interface ProductCreate {
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
}

export interface ProductUpdate extends ProductCreate {
  id: string
}

export interface CategoryCreate {
  name: string
}

export interface CouponCreate {
  promoCode: string
  amount: number
  unit: string
}

export interface CouponUpdate extends CouponCreate {
  id: string
  isActive: boolean
}
