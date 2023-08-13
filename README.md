<h1>Golang에서의 ElasticSearch 코드 작성</h1>

<h6>API를 연동해서, CRUD작업</h6>
- 메인 DB를 Mysql, Mongo가 아닌 ElasticSearch로 구성하여

<h6>해당 코드에서는 이론적인 내용은 다루지 않고, 코드로써는 어떻게 동작시키는지에 대해서 주롤 다룰 예정</h6>

- 부가적으로 메서드의 역할에 대해서는 다룰 수 있습니다.


## 레파지토리 생성 목적
이전에 작업했던 MySql, Mongo에서의 데이터 변화를 감지하는 모듈에 대한 변화를
ElasticSearch에 동기화 하는 로직을 구성한 바 있다.
```azure
MySql - https://github.com/04Akaps/mysql-elastic-event-cahcer-server-

Mongo - https://github.com/04Akaps/mongo-elastic-event-cahcer-server
```

이렇게 구성한 데이터를 이제 검색 엔진에서 어떻게 데이터를 가져오고 활용해야 하는지에 대한 코드르 작업하는 것을 목표로 삼고 있다.
- 거창하게 말했지 그냥, ElasticSearch 코드 작성하고 싶어서 생성한 레포

<h3>Docker</h3>
```
docker run -d -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e "ELASTIC_USERNAME=<사용할 이름>" \
-e "ELASTIC_PASSWORD=<사용할 패스워드>" \
--name elasticsearch-docker \
docker.elastic.co/elasticsearch/elasticsearch:7.14.0
```

