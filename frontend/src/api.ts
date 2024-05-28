import axios, {
  AxiosInstance,
  AxiosResponse,
  AxiosRequestConfig,
  InternalAxiosRequestConfig,
} from 'axios'
import {
  Credentials,
  ProductCreate,
  CategoryCreate,
  ProductUpdate,
  CouponCreate,
  CouponUpdate,
} from './pkg/requests'
import { ICategory, ICoupon, IOrder, IProduct } from './pkg/models'
import { ISignInResponse } from './auth/models'

interface IAPI {
  signUserIn(credentials: Credentials): Promise<AxiosResponse<any>>
  getProducts(id?: string, categories?: string): Promise<AxiosResponse<IProduct[]>>
  createProduct(product: ProductCreate): Promise<AxiosResponse<void>>
  updateProduct(product: ProductUpdate): Promise<AxiosResponse<void>>
  deleteProduct(id: string): Promise<AxiosResponse<string>>
  getCategories(): Promise<AxiosResponse<ICategory[]>>
  createCategory(category: CategoryCreate): Promise<AxiosResponse<string>>
  updateCategory(category: ICategory): Promise<AxiosResponse<string>>
  deleteCategory(id: string): Promise<AxiosResponse<void>>
  getCoupons(): Promise<AxiosResponse<ICoupon[]>>
  updateCoupon(coupon: CouponUpdate): Promise<AxiosResponse<void>>
  deleteCoupon(id: string): Promise<AxiosResponse<void>>
  createCoupon(coupon: CouponCreate): Promise<AxiosResponse<string>>
  getOrders(id?: string): Promise<AxiosResponse<IOrder[]>>
  updateOrder(order: IOrder): Promise<AxiosResponse<void>>
}

class API implements IAPI {
  private instance: AxiosInstance

  constructor(baseURL: string) {
    this.instance = axios.create({
      baseURL: baseURL,
      headers: { 'Content-Type': 'application/json' },
    })
    this.setupInterceptors()
    this.getCustomerId()
      .then((res: AxiosResponse<{ id: string }>) => res.data)
      .then((res) => {
        localStorage.setItem('customerId', res.id)
      })
      .catch((err: any) => {
        console.error(err)
      })
  }

  private customerIdInjector(
    requestConfig: InternalAxiosRequestConfig<any>,
  ): InternalAxiosRequestConfig<any> {
    const customerId: string = localStorage.getItem('customerId')
    if (requestConfig.method === 'get' || requestConfig.method === 'delete') {
      requestConfig.params = { ...requestConfig.params, customerId }
    } else {
      requestConfig.data = { ...requestConfig.data, customerId }
    }

    return requestConfig
  }

  private injectSessionToken<T>(
    requestFn: (config: AxiosRequestConfig) => Promise<AxiosResponse<T>>,
  ): () => Promise<AxiosResponse<T>> {
    const sessionToken: string = localStorage.getItem('sessionToken')
    const expiresAt: Date = new Date(localStorage.getItem('expiresAt'))
    const now = new Date()

    const timeDifference: number = expiresAt.getTime() - now.getTime()
    if (timeDifference < 10 * 60 * 1000) {
      this.injectSessionToken((mergedConfig) => this.instance.post('/user/refresh', mergedConfig))()
        .then((res) => res.data)
        .then((res: ISignInResponse) => {
          localStorage.setItem('sessionToken', res.sessionToken)
          now.setTime(now.getTime() + 60 * 60 * 1000)
          localStorage.setItem('expiresAt', now.toISOString())
        })
    }

    return async () => {
      const config: AxiosRequestConfig = {}
      if (sessionToken) {
        config.headers = {
          Authorization: `Bearer ${sessionToken}`,
        }
      }
      return requestFn(config)
    }
  }

  private async getCustomerId(): Promise<AxiosResponse<{ id: string }>> {
    // TODO: get tag based on URL
    const tag = 'dev'
    return this.instance.get(`/customer/id/${tag}`)
  }

  public async signUserIn(credentials: Credentials): Promise<AxiosResponse<any>> {
    return this.instance.post('/user/signin', { ...credentials })
  }

  public async getProducts(id?: string, categories?: string): Promise<AxiosResponse<IProduct[]>> {
    return this.instance.get('/product/', {
      params: {
        ...(id && { id }),
        ...(categories && { categories }),
      },
    })
  }

  public async createProduct(product: ProductCreate): Promise<AxiosResponse<void>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.post('/product/', product, mergedConfig),
    )()
  }

  public async updateProduct(product: ProductUpdate): Promise<AxiosResponse<void, any>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.put('/product/', product, mergedConfig),
    )()
  }

  public async deleteProduct(id: string): Promise<AxiosResponse<string>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.delete(`/product/${id}`, mergedConfig),
    )()
  }

  public async getCategories(): Promise<AxiosResponse<ICategory[]>> {
    return this.instance.get('/product/category')
  }

  public async createCategory(category: CategoryCreate): Promise<AxiosResponse<string>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.post('/product/category', category, mergedConfig),
    )()
  }

  public async updateCategory(category: ICategory): Promise<AxiosResponse<string>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.put('/product/category', category, mergedConfig),
    )()
  }

  public async deleteCategory(id: string): Promise<AxiosResponse<void>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.delete(`/product/category/${id}`, mergedConfig),
    )()
  }

  public async getCoupons(): Promise<AxiosResponse<ICoupon[]>> {
    return this.instance.get('/coupon')
  }

  public async updateCoupon(coupon: CouponUpdate): Promise<AxiosResponse<void>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.put('/coupon', coupon, mergedConfig),
    )()
  }

  public async deleteCoupon(id: string): Promise<AxiosResponse<void>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.delete(`/coupon/${id}`, mergedConfig),
    )()
  }

  public async createCoupon(coupon: CouponCreate): Promise<AxiosResponse<string>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.post('/coupon', coupon, mergedConfig),
    )()
  }

  public async getOrders(id?: string): Promise<AxiosResponse<IOrder[]>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.get('/order', {
        ...mergedConfig,
        ...(id && { id }),
      }),
    )()
  }

  public async updateOrder(order: IOrder): Promise<AxiosResponse<void>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.put('/order', order, mergedConfig),
    )()
  }

  public setupInterceptors() {
    this.instance.interceptors.request.use(
      (config) => this.customerIdInjector(config),
      (error) => Promise.reject(error),
    )
  }
}

// TODO: read URL from config
const api = new API('http://localhost:8000/api')
export default api
