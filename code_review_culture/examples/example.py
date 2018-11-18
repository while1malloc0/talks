def read_ips_from_file():
    with open(IP_FILE) as f:
        content = f.readlines()
    return content
