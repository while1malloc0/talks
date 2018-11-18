def read_ips_from_file():
    with open(IP_FILE) as f:
        content = f.readlines()
    return content

def process_ips():
    ipArray = read_ips_from_file()

    for ip in ipArray:
        process(ip)
