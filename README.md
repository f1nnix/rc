# rc

Simple and effective CLI-fronted for [rclone](https://github.com/ncw/rclone).

`rc` provides a flexible command subset for advanced bi-directional sync over multiple [rclone](https://github.com/ncw/rclone) remotes.

## How it works

Let's say we have several configured remotes in rclone config:

* `db` — Dropbox remote;
* `yd` — Yandex Disk remote;
* ...any other remotes with supported backend.

Assuming this, let's see how `rc` helps to manage vaults over these storage backends:

| Command                  | Description                                                                                 |
|:-------------------------|:--------------------------------------------------------------------------------------------|
| `rc db`                  | Sync local folder `~/db/` to remote `db:/`                                                  |
| `rc yd:/vim`             | Sync local folder `~/yd/vim/` to `yd:/vim/`                                                 |
| `rc docs`                | Encrypt local folder `~/Documents/work` and sync encrypted files to remote `db:/encrypted`, |
| `rc books:/Zelazny down` | Sync remote folder `db:/library/Zelazny/` to  `~/calibre/Zelazny`                             |

The core idea of `rc` is to define several higher-level remotes in rc config, which inherit actual `rclone` ones. When you run `rc books:/Zelazny down` command, the following pipeline is applied:

1. `rc` parses provided arguments:
    * `books` — remote name;
    * `/Zelazny` — so-called userPath;
    * `down` — sync direction. Down means "from remote to local", and vice versa.
2. `rc` searches for remote `books` in `.config.yml`. Let's say config is:
    ```yml
    base_dir: /Users/user
    options:
      exclude: [".DS_Store"]
    remotes:
      books:
        remote: db
        path: /library
        local_path: /calibre
    ```
3. `rc` build local dir path and remote one:
    1. **Local:** `base_dir + (remote | local_path) + userPath`. Will generate `/Users/user/calibre/Zelazny/`.
    2. **Remote:** `remote:remote_path + userPath`. Will generate `db:/library/Zelazny/`.
4. `rc` applies options from global config options, such as `exclude`.
5. `rc` runs `rclone --exclude=.DS_Store sync db:/library/Zelazny /Users/user/calibre/Zelazny`

### Encryption

You may wish to encrypt remote files and sync to remote directory like `enc_storage`. To do this:

1. Create with `rclone config` actual storage backend in rclone, for example `db<Dropbox>`, or use existing.
2. Create with `rclone config` crypt storage backend in rclone. For `path`, specify actual `remote` and `path`, where encrypted files will be physically placed on remote storage. Example: `db:/enc_storage`.
3. Create rc-remote in config, e. g. `docs`. For `remote` specify crypt storage from step 2.
    1. By default, the same `local_path` rules are applied: decrypted files locally will be placed to `<base_dir>/<remote_path>`. If you specify `local_path`, decrypted files will be stored in `<base_dir>/<local_path>`.
4. Run `rc docs`. It will encrypt `<base_dir>/<remote_path | local_path>` and upload to `<remote>:<path>`, specified in rclone config.

## Usage

`rc <remote:path> <direction>`

* `<remote:path>`: rc-remote and userPath to specify nested sync directory
* `<direction>`: `up` or `down`, sync direction. `up` is default and my be omited.

## Configuration

Example config:

```yaml
base_dir: /Users/user
remotes:
  stuff:
    remote: yd
  books:
    remote: yd
    path: /library
    local_path: /calibre
  docs:
    remote: crypt_docs
    remote_path: /
    local_path: /Documents/docs
```

`Config` params:

* `base_dir`: absolute path on filesystem, which all other remotes are relative to.
* `remotes`: a named list of Remotes.

`Remote` params:

* **`remote`, required**: rclone remote to use, should exist in rclone config;
* `remote_path`, optional: directory to sync on remote storage. Used as local_path, if the last is not provided.
* `local_path`, optional: directory to sync on local storage. *Overwrites* `remote_path` for building local target path.

## Known bugs and features

There are several bugs and featured, I'm going to implement in future. Please, feel free to open PR.

* `local_path` cannot contains spaces. For some reasons running `exec.Command()` with spaces in one of args raises incorrect argument parsing, even with quoting / escape. Running the same arguments in bash works properly.
* rc must failback to rclone remotes with the same argparse logic if requested remote was not found in rc config. Right now this code block is commented.

Please, note, **`rc` is in alpha.** Though critical functions are partly tested, use it at your own risk.

## License

Copyright (c) 2018 Ilya Rusanen

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.