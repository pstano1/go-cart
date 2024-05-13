export interface IProduct {
  id: string
  name: string
  categories: string[]
  descriptions: {
    [key: string]: string
  }
  prices: {
    [key: string]: number
  }
}
