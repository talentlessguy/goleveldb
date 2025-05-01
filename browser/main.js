import { fs, configure } from 'https://esm.sh/@zenfs/core@2.2.0'
import { InMemory } from 'https://esm.sh/@zenfs/core@2.2.0/backends/memory'

// Initialize the file system
await configure({
  mounts: {
    '/tmp': InMemory,
    '/home/user': InMemory,
  }
})

window.ZenFS = fs

await fs.promises.mkdir('/tmp/test')

const go = new Go()

go.env = {
  HOME: '/home/user',
  PATH: '/usr/bin:/usr/local/bin',
}

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
  go.run(result.instance)
})

console.log("Files in LevelDB:", await fs.promises.readdir('/tmp/test')) // Should include "test"
