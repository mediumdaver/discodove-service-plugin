/* Written by Dave Richards.
 *
 * Service interface definition.
 */
package discodove_interface_service

import (
	"net"
)

type DiscoDoveService interface {

	/* The real work is done here, this method will be called, in context of a go routine, for every 
	 * connection made on the listen port.  DiscoDove will not attempt to validate the connection in any
	 * way, the assumption is if it passed all the firewalls and discodove's controls (SSL for example),
	 * then it's over to you.
	 * 
	 * Please close the connection and exit when you're done.
	 *
	 * conn 			: the connection.
	 * authenticators	: a collection of authenticators to call in priority order.
	 * stores			: all stores that have been initiated for use.
	 * secure			: a method of encrypting a plain text session if needed, e.g. STARTTLS
	 */
    HandleConnection(net.Conn)
}