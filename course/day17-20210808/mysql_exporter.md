mysql 可用性
    操作失败
        select 1;
        ping
慢查询次数
    show global status where variable_name='slow_queries'
容量:
    qps:
        show global status where variable_name='Queries'
    tps:
        insert, update, delete *
        com_insert
        com_update
        com_delete
        com_select
        com_replace
    连接:
        show global status where variable_name='Threads_connected'
        show global variables where variable_name= 'max_connections';

    流量：
        show global status where variable_name='Bytes_received'
        show global status where variable_name='Bytes_sent'


// mysql连接信息 => mysql host, port, user