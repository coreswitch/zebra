/*
 * Bind socket to specified VRF. Use LD_PRELOAD to override socket() system call
 * then call SO_BINDTODEVICE for VRF binding. VRF name is get from environment
 * variable VRF.
 *
 *  $ sudo VRF=vrf-upstream LD_PRELOAD=./vrf_socket.so ./app
 *
 */
#define _GNU_SOURCE
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <net/if.h>
#include <unistd.h>
#include <dlfcn.h>

/* #define DEBUG */
#if defined DEBUG
#define debug(...) printf(__VA_ARGS__)
#else
#define debug(...)
#endif

typedef int (*socket_func_type)(int, int, int);

int socket(int domain, int type, int protocol)
{
  int ret;
  int sock;
  struct ifreq ifr;
  const char *vrfname = NULL;

  /* Find original socket function. */
  socket_func_type socket_orig;
  socket_orig = (socket_func_type)dlsym(RTLD_NEXT, "socket");

  /* Call original socket function. */
  sock =  socket_orig(domain, type, protocol);
  debug("VRF sock: %d\n", sock);
  if (sock < 0) {
    return sock;
  }

  /* Get VRF name from environment variable. */
  vrfname = getenv("VRF");
  debug("VRF name: %s\n", vrfname);
  if (vrfname == NULL) {
    return sock;
  }

  /* Bind the socket to the VRF. */
  memset(&ifr, 0, sizeof(struct ifreq));
  snprintf(ifr.ifr_name, sizeof(ifr.ifr_name), "%s", vrfname);
  ret = setsockopt(sock, SOL_SOCKET, SO_BINDTODEVICE, (void *)&ifr, sizeof(struct ifreq));
  debug("VRF bind return: %d\n", ret);
  if (ret < 0) {
    close(sock);
    return ret;
  }

  /* Return VRF binded socket. */
  return sock;
}
