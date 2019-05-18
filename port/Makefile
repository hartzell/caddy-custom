# $FreeBSD$

PORTNAME=	caddy-custom
APPNAME=caddy
DISTVERSIONPREFIX=	v
DISTVERSION=	0.0.1
PORTREVISION=	1
CATEGORIES=	www

MAINTAINER=	hartzell@alerce.com
COMMENT=	My custom caddy + plugins

USES=		go
USE_GITHUB=	yes
GH_ACCOUNT=	hartzell

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${PORTNAME}

PLIST_FILES=	bin/caddy etc/rc.d/caddy

do-build:
	@cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} build

do-install:
	${INSTALL_PROGRAM} ${WRKSRC}/${APPNAME} \
		${STAGEDIR}/${PREFIX}/bin/${APPNAME}
	${INSTALL_SCRIPT} dist/init/freebsd/${APPNAME} \
		${STAGEDIR}/${PREFIX}/etc/rc.d/${APPNAME}

.include <bsd.port.mk>
