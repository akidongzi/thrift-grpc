namespace go self_thrift

struct Req {
    1: string msg;
}

struct Res {
    1: string msg;
}

service SelfHelloThrift {
    Res SelfSayHi(1: Req req);
}

