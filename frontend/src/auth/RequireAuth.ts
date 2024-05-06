import m from 'mithril'

interface RequireAuthAttrs {
  allowedPermissions?: string[]
}

const RequireAuth: m.Component<RequireAuthAttrs> = {
  view: (vnode: m.Vnode<RequireAuthAttrs, any>): m.Children => {
    const allowedPermissions = vnode.attrs.allowedPermissions || []

    const sessionToken: string = localStorage.getItem('sessionToken')
    if (!sessionToken) {
      m.route.set('/signin')
    }

    // const userPermissons: string[] = localStorage.getItem('userPermissions').split(',') || []
    // if (
    //   allowedPermissions.length > 0 &&
    //   !allowedPermissions.some((permission) => userPermissons.includes(permission))
    // ) {
    //   m.route.set('/signin')
    // }

    return vnode.children
  },
}

export default RequireAuth
