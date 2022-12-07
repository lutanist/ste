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

