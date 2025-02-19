# tact-go2

# Build Requirements

## Windows

To build this project on Windows, you need to install GCC.

1. Follow the instructions at [MSYS2](https://www.msys2.org/) to install GCC.
2. Run the following command:
   ```sh
   pacman -S mingw-w64-ucrt-x86_64-gcc
   ```
3. This will install `gcc.exe` in `C:\msys64\ucrt64\bin`.
4. Add `C:\msys64\ucrt64\bin` to your `PATH` environment variable.
5. Now you can proceed with building the project.

### Enabling CGO in Go

To enable CGO when building with Go, set the `CGO_ENABLED` environment variable:

```sh
export CGO_ENABLED=1  # Unix-like OS
set CGO_ENABLED=1      # Windows (cmd)
$env:CGO_ENABLED=1     # Windows (PowerShell)
```

When building the project, use:

```sh
CGO_ENABLED=1 go build
```

On Windows, you may also need to specify the C compiler:

```sh
set CC=x86_64-w64-mingw32-gcc
CGO_ENABLED=1 go build
```

---

# 빌드 전 요구사항

## Windows

이 프로젝트를 Windows에서 빌드하려면 GCC를 설치해야 합니다.

1. [MSYS2](https://www.msys2.org/)의 안내를 따라 GCC를 설치합니다.
2. 다음 명령어를 실행합니다:
   ```sh
   pacman -S mingw-w64-ucrt-x86_64-gcc
   ```
3. 이 과정이 완료되면 `C:\msys64\ucrt64\bin` 폴더에 `gcc.exe`가 설치됩니다.
4. 환경 변수 `PATH`에 `C:\msys64\ucrt64\bin`을 추가합니다.
5. 이제 프로젝트를 빌드할 수 있습니다.

### Go에서 CGO 활성화

Go에서 CGO를 활성화하려면 `CGO_ENABLED` 환경 변수를 설정해야 합니다.

```sh
export CGO_ENABLED=1  # Unix 계열 OS
set CGO_ENABLED=1      # Windows (cmd)
$env:CGO_ENABLED=1     # Windows (PowerShell)
```

프로젝트 빌드시 다음 명령어를 사용합니다:

```sh
CGO_ENABLED=1 go build
```

Windows에서는 C 컴파일러를 지정해야 할 수도 있습니다:

```sh
set CC=x86_64-w64-mingw32-gcc
CGO_ENABLED=1 go build
```