// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Mytestlab123IochainTest4Iochaintest4 from './mytestlab123.iochain_test4.iochaintest4'


export default { 
  Mytestlab123IochainTest4Iochaintest4: load(Mytestlab123IochainTest4Iochaintest4, 'mytestlab123.iochain_test4.iochaintest4'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}