import m from 'mithril'
import axios, { AxiosResponse } from 'axios'
import { ISignInResponse } from '../auth/models'

interface ISignInView extends m.Component {
  isSent: boolean
  isSuccess: boolean
  handleSubmit: (event: Event) => void
}

const signIn: ISignInView = {
  handleSubmit: (event: Event): void => {
    event.preventDefault()

    const formData = new FormData(event.target as HTMLFormElement)
    let credentials: { [key: string]: string } = {}

    for (let [key, value] of formData.entries()) {
      credentials[key] = value as string
    }

    axios
      .post('http://localhost:8000/api/user/signin', {
        ...credentials,
      })
      .then((res: AxiosResponse<ISignInResponse>) => res.data)
      .then((data: ISignInResponse) => {
        localStorage.setItem('sessionToken', data.sessionToken)
        // localStorage.setItem('permissons', data.permissions.join(','))
      })
      .then(() => {
        m.route.set('/')
      })
      .catch((err: any) => {
        // pass
      })
  },
  view: () => {
    return m(
      'main',
      { class: 'flex justify-center items-center h-screen bg-midnightGreen' },
      m(
        'form',
        {
          name: 'signInForm',
          onsubmit: (event: Event) => signIn.handleSubmit(event),
          class: 'bg-antiflashWhite p-16 rounded-lg shadow-md',
        },
        m('h1', { class: 'uppercase text-lg font-semibold text-midnightGreen' }, 'Sign In'),
        m('input', {
          class:
            'w-full p-2 mt-2 mb-2 focus:outline-none bg-antiflashWhite border-b-2 border-black focus:border-midnightGreen',
          type: 'text',
          name: 'username',
        }),
        m('input', {
          class:
            'w-full p-2 mt-2 mb-2 focus:outline-none bg-antiflashWhite border-b-2 border-black focus:border-midnightGreen',
          type: 'password',
          name: 'password',
        }),
        m(
          'button',
          {
            type: 'submit',
            class: 'bg-midnightGreen text-antiflashWhite w-full rounded-lg p-2 mt-2 uppercase',
          },
          'sign in',
        ),
      ),
    )
  },
  isSent: false,
  isSuccess: false,
}

export default signIn
