import axios, {
  AxiosInstance,
  AxiosResponse,
  AxiosRequestConfig,
  InternalAxiosRequestConfig,
} from 'axios'
import { Credentials, ProductCreate, CategoryCreate, ProductUpdate } from './pkg/requests'
import { IProduct } from './pkg/models'

interface IAPI {
  signUserIn(credentials: Credentials): Promise<AxiosResponse<any>>
  getProducts(id?: string, categories?: string): Promise<AxiosResponse<IProduct[]>>
  createProduct(product: ProductCreate): Promise<AxiosResponse<void>>
  updateProduct(product: ProductUpdate): Promise<AxiosResponse<void>>
  deleteProduct(id: string): Promise<AxiosResponse<string>>
  getCategories(): Promise<AxiosResponse<string[]>>
  createCategory(category: CategoryCreate): Promise<AxiosResponse<string>>
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

  public async getCategories(): Promise<AxiosResponse<string[]>> {
    return this.instance.get('/product/category')
  }

  public async createCategory(category: CategoryCreate): Promise<AxiosResponse<string>> {
    return this.injectSessionToken((mergedConfig) =>
      this.instance.post('/product/category', category, mergedConfig),
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