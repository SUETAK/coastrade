FROM mysql
#ポートを開ける
EXPOSE 3306

#MySQL設定ファイルをイメージ内にコピー
COPY ./build/docker/mysql /etc/mysql/conf.d/my.cnf

#docker runに実行される
CMD ["mysqld"]