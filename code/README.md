# module8/code

1、docker build -t gaojingwen/module8:v9 --platform linux/amd64 .
2、docker push gaojingwen/module8:v9
3、kubectl create ns module
4、kubectl apply -f configmap.yaml
5、kubectl apply -f deployment.yaml
6、kubectl apply -f service.yaml
7、openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=amber.com/O=amber"
8、kubectl create secret tls module8-tls -n moudle --cert=./tls.crt --key=./tls.key
9、kubectl apply -f ingress.yaml
10、curl amber.com/healthz
