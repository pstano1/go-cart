export interface Credentials {
  customerId?: string
  username: string
  password: string
}

export interface ProductCreate {
  name: string
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
