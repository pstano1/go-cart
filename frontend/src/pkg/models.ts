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
