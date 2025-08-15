interface FetchesDictionary {
    [key: string] : Array<() => Promise<any>>
}

interface LocksDictionary {
    [key: string] : boolean
}

class FetchQueue {
    fetches: FetchesDictionary = {}
    locks: LocksDictionary = {}
    
    constructor() {}

    addFetch(name: string, fetch: () => Promise<any>) {
        if (this.fetches[name] == undefined) {
            this.fetches[name] = []
        }
        if (this.locks[name] == undefined) {
            this.locks[name] = false
        }
        if (this.fetches[name].length !== 0) {
            return
        }
        this.fetches[name].push(fetch)
        this.execute(name)
    }

    execute(name: string) {
        if (this.locks[name]) return
        this.locks[name] = true
        console.log(this.fetches[name][0])
        this.fetches[name][0]().then(() => {this.locks[name] = false}, () => {this.locks[name] = false})
    }
}

export const Queue = new FetchQueue()