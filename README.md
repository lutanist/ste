# ste
Security Token Exchanger. [부실채권(NPL)](http://koreaifp.com/financial_magazine/501) 뿐만 아니라 증권토큰 거래를 꿈꾸며 작명

## 정의
[시큐리티 토큰(Security Token), 왜 하는 건데?](https://medium.com/@ksjterry/%EC%8B%9C%ED%81%90%EB%A6%AC%ED%8B%B0-%ED%86%A0%ED%81%B0-%EC%95%88%EB%82%B4%EC%84%9C-54a5632dbb60)

## 인증

### OAuth, OIDC
[zitadel/oidc github](https://github.com/zitadel/oidc)
https://developer.okta.com/docs/guides/sign-into-web-app-redirect/go/main/
[OpenID Certified](https://openid.net/developers/certified/)
[Ory](https://www.ory.sh/docs/ecosystem/projects)
[Amazon Coginito](https://docs.aws.amazon.com/ko_kr/cognito/latest/developerguide/what-is-amazon-cognito.html)
[go-jose](https://github.com/go-jose/go-jose) golang JWT implementation

### 1원 계좌 인증
1원 계좌 인증 api 제공업체들이 꽤 있음. [오픈뱅킹](https://developers.kftc.or.kr/dev/openapi/open-banking/oauth)을 사용하는건가? 아니면 업체를 사용하나?

## 원장

https://github.com/hyperledger/fabric


## Backend 구현
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [golang tool dependencies](https://play-with-go.dev/tools-as-dependencies_go119_en/)
- [ent](https://entgo.io/docs/getting-started/)
- golang configuration. [viper](https://github.com/spf13/viper)
  - [merge config example](https://golang.hotexamples.com/examples/github.com.spf13.viper/Viper/MergeConfig/golang-viper-mergeconfig-method-examples.html)
  - vipder flag 자동으로 되는지 확인
- [fiber docs](https://docs.gofiber.io/)
  - https://blog.gopheracademy.com/advent-2014/configuration-with-fangs/
- [golang fiber fullstack example](https://github.com/divrhino/divrhino-trivia-fullstack)
- [위에거 text 링킁](https://divrhino.com/articles/full-stack-go-fiber-with-docker-postgres/)
- [fiber-go-template](https://github.com/create-go-app/fiber-go-template)

https://github.com/alpody/golang-fiber-realworld-example-app

https://dev.to/komfysach/go-live-reload-using-air-40ll

https://github.com/gofiber/boilerplate/blob/master/app.go

### wire

https://www.shipyardapp.com/blog/go-dependency-injection-wire/

https://syntaxsugar.tistory.com/entry/Golang-Dependency-Injection

https://syntaxsugar.tistory.com/entry/koWire-Jacket-IoC-Container-of-googlewire-for-cloud-native

settings.json에 아래 내용 추가
```json
{
    "gopls": {
        "build.buildFlags": ["-tags=wireinject"]
    }
}
```

또 다른 di famework https://github.com/uber-go/dig

### https 개발환경

https://web.dev/how-to-use-local-https/

https://github.com/FiloSottile/mkcert

```winget install FiloSottile.mkcert```

C:\Users\lutan\AppData\Local\Microsoft\WinGet\Links 에 설치됨

### cli

이것도 괜찮은 듯
https://cli.urfave.org/v2/getting-started/ 

### patterns

https://velog.io/@tae2089/Go%EC%97%90%EC%84%9C-Builder-Pattern-%EC%82%AC%EC%9A%A9%ED%95%B4%EB%B3%B4%EA%B8%B0

https://www.sohamkamani.com/golang/options-pattern/

https://www.sohamkamani.com/golang/2018-06-20-golang-factory-patterns/

oidc

https://ssup2.github.io/programming/Golang_Google_OIDC_%EC%9D%B4%EC%9A%A9/

gotemplate
- https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax
- https://pkg.go.dev/text/template#hdr-Functions

singleton, once & init func
https://stackoverflow.com/questions/67334017/having-a-singleton-pattern-in-go-wire-injection