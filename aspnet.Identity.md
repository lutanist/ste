```mermaid
classDiagram
    direction BT

    class AuthenticationScheme {
        String Name()
        String DisplayName()
        Type HandlerType()
    }

    class AuthenticationProperties {

    }

    class AuthenticationSchemeOptions

    class AuthenticateResult

    %% Created per request to handle authentication for a particular scheme
    class IAuthenticationHandler {
        <<interface>>
        InitializeAsync(AuthenticationScheme schema, HttpContext context)
        AuthenticateAsync() Task~AuthenticateResult~
        ChallengeAsync(AuthenticationProperties? properties)
        ForbidAsync(AuthenticationProperties? properties)
    }
    IAuthenticationHandler ..> AuthenticateResult
    IAuthenticationHandler ..> AuthenticationProperties
    IAuthenticationHandler ..> AuthenticationScheme
   
    class AuthenticationHandler~TOptions~ {
        <<Abstract>>
        #InitializeHandlerAsync()
        #HandleAuthenticateAsync()*
        #HandleForbiddenAsync()
        #HandleChallengeAsync()
    }
    AuthenticationHandler~TOptions~ ..|> IAuthenticationHandler
    AuthenticationHandler~TOptions~ --* AuthenticationScheme: Schema

    %% Used to determine if a handler supports SignOut.
    class IAuthenticationSignOutHandler {
        <<interface>>
        SignOutAsync(AuthenticationProperties? properties)
    }
    IAuthenticationSignOutHandler --|> IAuthenticationHandler

    %% Used to determine if a handler supports SignIn.
    class IAuthenticationSignInHandler {
        <<interface>>
        SignInAsync(ClaimsPrincipal user, AuthenticationProperties? properties);
    }
    IAuthenticationSignInHandler --|> IAuthenticationSignOutHandler

    class SignOutAuthenticationHandler~TOptions~ {
        <<Abstract>>
        HandleSignOutAsync(AuthenticationProperties? properties)*
    }
    SignOutAuthenticationHandler~TOptions~ --|> AuthenticationHandler~TOptions~
    SignOutAuthenticationHandler~TOptions~ --|> IAuthenticationSignOutHandler

    class SignInAuthenticationHandler~TOptions~ {
        <<Abstract>>
        HandleSignInAsync(ClaimsPrincipal user, AuthenticationProperties? properties)
    }
    SignInAuthenticationHandler~TOptions~ --|> SignOutAuthenticationHandler~TOptions~
    SignInAuthenticationHandler~TOptions~ --|> IAuthenticationSignInHandler

    class CookieAuthenticationHandler {
        InitializeHandlerAsync()
    }
    CookieAuthenticationHandler --|> SignInAuthenticationHandler~CookieAuthenticationOptions~

    class IAuthenticationRequestHandler {
        <<interface>>
        HandleRequestAsync() bool
    }
    IAuthenticationRequestHandler --|>IAuthenticationHandler

    class RemoteAuthenticationOptions
    RemoteAuthenticationOptions --|> AuthenticationSchemeOptions

    class RemoteAuthenticationHandler~TOptions~ {
        <<Abstract>>
    }
    RemoteAuthenticationHandler~TOptions~ --|> AuthenticationHandler~TOptions~
    RemoteAuthenticationHandler~TOptions~ --|> IAuthenticationRequestHandler


    class OAuthOptions
    OAuthOptions --|> RemoteAuthenticationOptions

    class OAuthHandler~TOptions~
    OAuthHandler~TOptions~ --|> RemoteAuthenticationHandler~TOptions~

    class GoogleHandler
    GoogleHandler --|> OAuthHandler~GoogleOptions~

    class FacebookHandler
    FacebookHandler --|> OAuthHandler~FacebookOptions~

    class TwitterHandler
    TwitterHandler --|> RemoteAuthenticationHandler~TwitterOptions~

    class JwtBearerHandler
    JwtBearerHandler --|> AuthenticationHandler~JwtBearerOptions~

```

- `AuthenticationProperties` Dictionary used to store state values about the authentication session.

### AuthenticationBuilder

```mermaid
classDiagram
    direction LR

    class AuthenticationOptions {
        IList<AuthenticationSchemeBuilder> _schemes
        IEnumerable<AuthenticationSchemeBuilder> Schemes
        IDictionary<string, AuthenticationSchemeBuilder> SchemeMap
        string? DefaultScheme
        string? DefaultAuthenticateScheme
        string? DefaultSignInScheme
        string? DefaultSignOutScheme
        string? DefaultChallengeScheme
        bool RequireAuthenticatedSignIn

        AddScheme(string name, Action~AuthenticationSchemeBuilder~ configureBuilder)
    }

    class AuthenticationSchemeOptions {
        string? ClaimsIssuer
        object? Events
        Type? EventsType
        string? ForwardDefault
        string? ForwardAuthenticate
        string? ForwardChallenge
        string? ForwardForbid
        string? ForwardSignIn
        string? ForwardSignOut

        Validate() void
    }

    class AuthenticationBuilder {
        IServiceCollection Services

        AddSchemeHelper(string authenticationScheme, string? displayName, Action~TOptions~? configureOptions) AuthenticationBuilder

        AddSchema() AuthenticationBuilder

        AddRemoteScheme() AuthenticationBuilder
    }
    AuthenticationBuilder ..|> AuthenticationOptions

    class CookieAuthenticationOptions {
        CookieBuilder Cookie
        IDataProtectionProvider? DataProtectionProvide
        bool SlidingExpiration
        PathString LoginPath
        PathString AccessDeniedPath
        string ReturnUrlParameter
        CookieAuthenticationEvents Events
        ISecureDataFormat<AuthenticationTicket> TicketDataFormat
        ICookieManager CookieManager
        ITicketStore? SessionStore
        TimeSpan ExpireTimeSpan
    }
    CookieAuthenticationOptions --|> AuthenticationSchemeOptions

    class CookieExtensions {
        <<Extension>>
        AddCookie(string authenticationScheme, string? displayName, Action~CookieAuthenticationOptions~ configureOptions) AuthenticationBuilder
    }
    CookieExtensions ..|> AuthenticationBuilder


    class JwtBearerOptions {
        List~ISecurityTokenValidator~ SecurityTokenValidators
        bool RequireHttpsMetadata
        string MetadataAddress
        string? Authority
        string? Audience
        string Challenge
        JwtBearerEvents Events
        HttpMessageHandler? BackchannelHttpHandler
        HttpClient Backchannel
        TimeSpan BackchannelTimeout
        OpenIdConnectConfiguration? Configuration
        IConfigurationManager~OpenIdConnectConfiguration~? ConfigurationManage
        bool RefreshOnIssuerKeyNotFound
        TokenValidationParameters TokenValidationParameters
        bool SaveToken
        bool IncludeErrorDetails
        bool MapInboundClaims
        TimeSpan AutomaticRefreshInterval
        TimeSpan RefreshInterval
    }
    JwtBearerOptions --|> AuthenticationSchemeOptions

    class JwtBearerExtensions {
        <<Extension>>
        AddJwtBearer() AuthenticationBuilder
    }
    JwtBearerExtensions ..|> AuthenticationBuilder


    class OAuthExtensions {
        <<Extension>>
        AddOAuth() AuthenticationBuilder
    }
    OAuthExtensions ..|> AuthenticationBuilder


    class IApplicationBuilder {
        <<Interface>>
        UseMiddleware() IApplicationBuilder
    }

    class AuthenticationMiddleware

    class AuthAppBuilderExtensions {
        <<Extension>>
        UseAuthentication() IApplicationBuilder
    }
    AuthAppBuilderExtensions ..|> IApplicationBuilder
    AuthAppBuilderExtensions ..|> AuthenticationMiddleware

    class IServiceCollection {
        <<Interface>>
    }

    class AuthenticationService

    class NoopClaimsTransformation

    class AuthenticationHandlerProvider

    class AuthenticationCoreServiceCollectionExtensions {
        <<Extension>>
        AddAuthenticationCore() IServiceCollection
    }
    AuthenticationCoreServiceCollectionExtensions ..|> IServiceCollection
    AuthenticationCoreServiceCollectionExtensions ..|> AuthenticationService
    AuthenticationCoreServiceCollectionExtensions ..|> NoopClaimsTransformation
    AuthenticationCoreServiceCollectionExtensions ..|> AuthenticationHandlerProvider
    AuthenticationCoreServiceCollectionExtensions ..|> AuthenticationSchemeProvider

    class AuthenticationServiceCollectionExtensions {
        <<Extension>>
        AddAuthentication() IServiceCollection
    }
    AuthenticationCoreServiceCollectionExtensions ..|> IServiceCollection
    AuthenticationCoreServiceCollectionExtensions ..|> AuthenticationCoreServiceCollectionExtensions

```

### AddAuthentication 
```mermaid
sequenceDiagram
    app ->> AuthenticationServiceCollectionExtensions: AddAuthentication()
    AuthenticationServiceCollectionExtensions ->> AuthenticationCoreServiceCollectionExtensions: AddAuthenticationCore()
        AuthenticationCoreServiceCollectionExtensions ->> IServiceCollection: TryAddScoped<IAuthenticationService, AuthenticationService>()
        AuthenticationCoreServiceCollectionExtensions ->> IServiceCollection: TryAddSingleton<IClaimsTransformation, NoopClaimsTransformation>()
        AuthenticationCoreServiceCollectionExtensions ->> IServiceCollection: TryAddScoped<IAuthenticationHandlerProvider, AuthenticationHandlerProvider>()
        AuthenticationCoreServiceCollectionExtensions ->> IServiceCollection: TryAddSingleton<IAuthenticationSchemeProvider, AuthenticationSchemeProvider>()
        AuthenticationCoreServiceCollectionExtensions --) AuthenticationServiceCollectionExtensions: IServiceCollection
    AuthenticationServiceCollectionExtensions ->> IServiceCollection: AddDataProtection()
    AuthenticationServiceCollectionExtensions ->> IServiceCollection: AddWebEncoders()
    AuthenticationServiceCollectionExtensions ->> IServiceCollection: TryAddSingleton<ISystemClock, SystemClock>()
    AuthenticationServiceCollectionExtensions --) app: AuthenticationBuilder
```