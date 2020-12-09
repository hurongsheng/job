docker run -it --name laravel -p 81:80 -v /Users/hurongsheng/work/project/php/phpcms:/var/www/html/app hurongsheng/laravelsys
docker run -it --name laravel -p 81:80 -v /Users/hurongsheng/work/project/php/phpcms:/var/www/html/app 6aa1cb97e6cb

docker exec  -it laravel /bin/bash

docker exec  -it laravelsys /bin/bash
docker commit {{image-id}} hurongsheng/laravelsys:latest
docker push hurongsheng/laravelsys
