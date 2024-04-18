namespace go hello_thrift

struct Req {
    1: string msg;
}

struct Res {
    1: string msg;
}

service HelloThrift {
    Res SayHi(1: Req req);
}

